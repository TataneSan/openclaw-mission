const axios = require('axios');

const LLM_BASE_URL = process.env.LLM_BASE_URL || 'https://llm.contes.me';
const LLM_API_KEY = process.env.LLM_API_KEY || '';

// Available models mapping
const MODEL_MAP = {
  'deepseek-v4-pro': 'deepseek-v4-pro',
  'deepseek-v4-flash': 'deepseek-v4-flash',
  'kimi-k2.6': 'kimi-k2.6',
  'minimax-m2.7': 'minimax-m2.7',
  'mimo-v2.5-pro': 'mimo-v2.5-pro',
  'skyclaw-v1': 'skyclaw-v1',
  // Aliases
  'gpt-4': 'deepseek-v4-pro',
  'gpt-3.5-turbo': 'deepseek-v4-flash',
  'claude-3-opus': 'deepseek-v4-pro',
  'claude-3-sonnet': 'kimi-k2.6',
};

function resolveModel(requestedModel) {
  return MODEL_MAP[requestedModel] || requestedModel;
}

async function proxyChatCompletion(requestBody, apiKey) {
  const startTime = Date.now();

  // Resolve model
  const resolvedModel = resolveModel(requestBody.model);

  // Build upstream request
  const upstreamRequest = {
    ...requestBody,
    model: resolvedModel,
    stream: false // We handle streaming separately
  };

  // Remove any client-specific fields
  delete upstreamRequest.api_key;

  try {
    const response = await axios.post(`${LLM_BASE_URL}/v1/chat/completions`, upstreamRequest, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${LLM_API_KEY}`,
      },
      timeout: 120000, // 2 minute timeout
    });

    const latencyMs = Date.now() - startTime;
    const usage = response.data.usage || { prompt_tokens: 0, completion_tokens: 0 };

    return {
      success: true,
      data: response.data,
      usage: {
        prompt_tokens: usage.prompt_tokens || 0,
        completion_tokens: usage.completion_tokens || 0,
        total_tokens: (usage.prompt_tokens || 0) + (usage.completion_tokens || 0)
      },
      latencyMs,
      resolvedModel
    };
  } catch (error) {
    const latencyMs = Date.now() - startTime;

    if (error.response) {
      return {
        success: false,
        error: error.response.data,
        statusCode: error.response.status,
        latencyMs,
        resolvedModel
      };
    }

    return {
      success: false,
      error: { message: error.message },
      statusCode: 500,
      latencyMs,
      resolvedModel
    };
  }
}

async function proxyStreamCompletion(requestBody, apiKey, res) {
  const startTime = Date.now();
  const resolvedModel = resolveModel(requestBody.model);

  const upstreamRequest = {
    ...requestBody,
    model: resolvedModel,
    stream: true
  };
  delete upstreamRequest.api_key;

  try {
    const response = await axios.post(`${LLM_BASE_URL}/v1/chat/completions`, upstreamRequest, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${LLM_API_KEY}`,
      },
      timeout: 120000,
      responseType: 'stream'
    });

    // Set SSE headers
    res.setHeader('Content-Type', 'text/event-stream');
    res.setHeader('Cache-Control', 'no-cache');
    res.setHeader('Connection', 'keep-alive');

    let totalTokens = 0;
    let completionTokens = 0;

    response.data.on('data', (chunk) => {
      const lines = chunk.toString().split('\n');
      for (const line of lines) {
        if (line.startsWith('data: ')) {
          const data = line.slice(6);
          if (data === '[DONE]') {
            res.write('data: [DONE]\n\n');
            return;
          }
          try {
            const parsed = JSON.parse(data);
            if (parsed.usage) {
              totalTokens = parsed.usage.total_tokens || 0;
              completionTokens = parsed.usage.completion_tokens || 0;
            }
            res.write(`data: ${JSON.stringify(parsed)}\n\n`);
          } catch (e) {
            // Skip invalid JSON
          }
        }
      }
    });

    response.data.on('end', () => {
      const latencyMs = Date.now() - startTime;
      res.end();

      // Log usage asynchronously
      return {
        success: true,
        usage: {
          prompt_tokens: totalTokens - completionTokens,
          completion_tokens: completionTokens,
          total_tokens: totalTokens
        },
        latencyMs,
        resolvedModel
      };
    });

    response.data.on('error', (err) => {
      res.end();
      return {
        success: false,
        error: { message: err.message },
        statusCode: 500,
        latencyMs: Date.now() - startTime,
        resolvedModel
      };
    });

  } catch (error) {
    const latencyMs = Date.now() - startTime;
    return {
      success: false,
      error: error.response?.data || { message: error.message },
      statusCode: error.response?.status || 500,
      latencyMs,
      resolvedModel
    };
  }
}

function listAvailableModels() {
  return [
    {
      id: 'deepseek-v4-pro',
      name: 'DeepSeek V4 Pro',
      description: 'Most capable model for complex reasoning',
      pricing: { prompt: 50, completion: 150 }
    },
    {
      id: 'deepseek-v4-flash',
      name: 'DeepSeek V4 Flash',
      description: 'Fast and efficient for everyday tasks',
      pricing: { prompt: 20, completion: 60 }
    },
    {
      id: 'kimi-k2.6',
      name: 'Kimi K2.6',
      description: 'Balanced performance and speed',
      pricing: { prompt: 40, completion: 120 }
    },
    {
      id: 'minimax-m2.7',
      name: 'MiniMax M2.7',
      description: 'Creative and conversational',
      pricing: { prompt: 30, completion: 90 }
    },
    {
      id: 'mimo-v2.5-pro',
      name: 'MiMo V2.5 Pro',
      description: 'Advanced reasoning and coding',
      pricing: { prompt: 45, completion: 135 }
    },
    {
      id: 'skyclaw-v1',
      name: 'SkyClaw V1',
      description: 'Specialized for code generation',
      pricing: { prompt: 35, completion: 105 }
    }
  ];
}

module.exports = {
  proxyChatCompletion,
  proxyStreamCompletion,
  listAvailableModels,
  resolveModel
};
