export class LoginSuccess {
  name: string;
  token: string;
  expires: Date;
  path: string;
  domain: string;
  secure: boolean;
  httpOnly: boolean;

  constructor(
    name: string,
    token: string,
    expires: number,
    path: string,
    domain: string,
    secure: boolean,
    httpOnly: boolean
  ) {
    this.name = name;
    this.token = token;
    this.expires = new Date(expires);
    this.path = path;
    this.domain = domain;
    this.secure = secure;
    this.httpOnly = httpOnly;
  }
}

export class LoginFailure {
  error: string;
  constructor(error: string) {
    this.error = error;
  }
}
