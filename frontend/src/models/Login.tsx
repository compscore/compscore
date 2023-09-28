export type LoginSuccess = {
  name: string;
  token: string;
  expires: Date;
  path: string;
  domain: string;
  secure: boolean;
  httpOnly: boolean;
};

export type LoginFailure = {
  error: string;
};
