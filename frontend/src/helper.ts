import { BootstrapSwal } from "components/BootstrapSwal";
import { onPromiseRejected } from "pages/ErrorPage";
import { NavigateFunction } from "react-router-dom";
import Swal from "sweetalert2";

//==============================================================================
// Promise
//==============================================================================

const DelayedPromise = <T>(value: T | Promise<T>, ms: number = 1000) =>
  value instanceof Promise
    ? delayPromise(value, ms)
    : new Promise<T>((resolve) => setTimeout(() => resolve(value), ms));

const delayPromise = async <T>(promise: Promise<T>, ms: number) => {
  await promise;
  await DelayedPromise(0, ms);
  return promise;
};

//==============================================================================
// Array
//==============================================================================

const arrayFromSize = <T>(
  size: number,
  elementGenerate: (index: number) => T
) => Array.from(Array(size).keys(), (_, index) => elementGenerate(index));

//==============================================================================
// Form
//==============================================================================

const formValidationClassName = (error: any) =>
  error !== undefined ? "is-invalid" : "";

//==============================================================================
// MetaMask
//==============================================================================

const userRejectTransaction = (reason: any) =>
  (reason as { code: number }).code !== undefined && reason.code === 4001;

const handleRejectMetaMaskPromise = (
  reason: any,
  navigate: NavigateFunction
) => {
  Swal.close();
  if (userRejectTransaction(reason)) {
    BootstrapSwal.fire({
      icon: "error",
      title: "You rejected the transaction.",
    });
  } else {
    onPromiseRejected(reason, navigate);
  }
};

//==============================================================================
// Date
//==============================================================================

const toDDMMYYYYstring = (date: Date) =>
  `${toTwoDigits(date.getDate())}/${toTwoDigits(
    date.getMonth() + 1
  )}/${date.getFullYear()}`;

const toTwoDigits = (n: number) => (n >= 10 ? n.toString() : `0${n}`);

const dateFromDDMMYYYY = (date: string, separator: string = "/") => {
  const ddmmyyyy = date.split(separator).map((value) => parseInt(value));
  return new Date(ddmmyyyy[2], ddmmyyyy[1] - 1, ddmmyyyy[0]);
};

const dateFromYYYYMMDD = (date: string, separator: string = "/") => {
  const ddmmyyyy = date.split(separator).map((value) => parseInt(value));
  return new Date(ddmmyyyy[0], ddmmyyyy[1] - 1, ddmmyyyy[2]);
};

export {
  DelayedPromise,
  arrayFromSize,
  formValidationClassName,
  userRejectTransaction,
  handleRejectMetaMaskPromise,
  toDDMMYYYYstring,
  dateFromDDMMYYYY,
  dateFromYYYYMMDD,
};
