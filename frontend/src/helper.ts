const DelayedPromise = <T>(value: T | Promise<T>, ms: number = 1000) =>
  value instanceof Promise
    ? delayPromise(value, ms)
    : new Promise<T>((resolve) => setTimeout(() => resolve(value), ms));

const delayPromise = async <T>(promise: Promise<T>, ms: number) => {
  await promise;
  await DelayedPromise(0, ms);
  return promise;
};

const arrayFromSize = <T>(
  size: number,
  elementGenerate: (index: number) => T
) => Array.from(Array(size).keys(), (_, index) => elementGenerate(index));

const formValidationClassName = (error: any) =>
  error !== undefined ? "is-invalid" : "";

const userRejectTransaction = (reason: any) =>
  "code" in reason && reason.code === 4001;

export {
  DelayedPromise,
  arrayFromSize,
  formValidationClassName,
  userRejectTransaction,
};
