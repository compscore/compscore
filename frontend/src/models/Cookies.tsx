import { CookieSetOptions } from "universal-cookie";

export type cookies = {
  auth?: any;
  admin?: any;
};

export type setCookie = (
  name: "auth" | "admin",
  value: any,
  options?: CookieSetOptions | undefined
) => void;

export type removeCookie = (
  name: "auth" | "admin",
  options?: CookieSetOptions | undefined
) => void;
