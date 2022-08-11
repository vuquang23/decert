declare module "*.svg" {
  export const ReactComponent: any;
}

declare module "*.jpg" {
  const value: any;
  export default value;
}

declare module "crypto-browserify" {
  export namespace createHash {
    type TypedArray =
      | Uint8Array
      | Uint8ClampedArray
      | Uint16Array
      | Uint32Array
      | Int8Array
      | Int16Array
      | Int32Array
      | Float32Array
      | Float64Array;

    interface HashAlgorithm {
      digest(target: encoding): string;
      digest(): Buffer;

      update(
        data: string | Buffer | TypedArray | DataView,
        encoding?: string
      ): this;
      write(
        data: string | Buffer | TypedArray | DataView,
        encoding?: string
      ): this;

      end(): void;
      read(): void;
    }

    type encoding = "utf8" | "hex" | "base64";
    type algorithm =
      | "md5"
      | "rmd160"
      | "ripemd160"
      | "sha"
      | "sha1"
      | "sha224"
      | "sha256"
      | "sha384"
      | "sha512";
  }

  export function createHash(
    algorithm: createHash.algorithm,
    options?: any
  ): createHash.HashAlgorithm;
}
