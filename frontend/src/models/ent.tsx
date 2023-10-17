export type Check = {
  name: string;
  edges: {
    status?: Status;
    credential?: Credential;
  };
};

export type Status = {
  error: string;
  status: string;
  time: string;
  edges: {
    check?: Check;
    team?: Team;
    round?: Round;
  };
};

export type Credential = {
  password: string;
  edges: {
    check?: Check;
    team?: Team;
  };
};

export type Team = {
  number: number;
  name: string;
  role: string;
  edges: {
    status?: Status;
    credential?: Credential;
    round?: Round;
  };
};

export type Round = {
  number: number;
  edges: {
    status?: Status;
    team?: Team;
  };
};
