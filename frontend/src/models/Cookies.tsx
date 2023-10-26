import { CookieSetOptions } from "universal-cookie";

export type cookies = {
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  auth?: any;
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  admin?: any;
};

export type setCookie = (
  name: "auth" | "admin",
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  value: any,
  options?: CookieSetOptions | undefined
) => void;

export type removeCookie = (
  name: "auth" | "admin",
  options?: CookieSetOptions | undefined
) => void;
