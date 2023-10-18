export type LoginSuccess = {
  name: string;
  token: string;
  expiration: number;
  path: string;
  domain: string;
  secure: boolean;
  httpOnly: boolean;
};

export type LoginFailure = {
  error: string;
};
