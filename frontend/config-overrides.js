module.exports = function override(config, _env) {
  config.resolve.fallback = { stream: require.resolve("stream-browserify") };
  return config;
};
