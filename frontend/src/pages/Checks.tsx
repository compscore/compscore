import { CookieSetOptions } from "universal-cookie";

type Props = {
  setCookie: (
    name: "auth",
    value: any,
    options?: CookieSetOptions | undefined
  ) => void;
};

export default function Login({ setCookie }: Props) {
  return <></>;
}
