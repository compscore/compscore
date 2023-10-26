import { CookieSetOptions } from "universal-cookie";

export type cookies = {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  auth?: any;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  admin?: any;
};

export type setCookie = (
  name: "auth" | "admin",
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  value: any,
  options?: CookieSetOptions | undefined
) => void;

export type removeCookie = (
  name: "auth" | "admin",
  options?: CookieSetOptions | undefined
) => void;
