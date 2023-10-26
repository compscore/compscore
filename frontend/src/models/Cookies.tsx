import { CookieSetOptions } from "universal-cookie";

export type cookies = {
  // @ts-ignore
  auth?: any;
  // @ts-ignore
  admin?: any;
};

export type setCookie = (
  name: "auth" | "admin",
  // @ts-ignore
  value: any,
  options?: CookieSetOptions | undefined
) => void;

export type removeCookie = (
  name: "auth" | "admin",
  options?: CookieSetOptions | undefined
) => void;
