const DelayedPromise = <T>(value: T | Promise<T>, ms: number = 1000) =>
  value instanceof Promise
    ? delayPromise(value, ms)
    : new Promise<T>((resolve) => setTimeout(() => resolve(value), ms));

const delayPromise = async <T>(promise: Promise<T>, ms: number) => {
  await promise;
  await DelayedPromise(0, ms);
  return promise;
};

export { DelayedPromise };
