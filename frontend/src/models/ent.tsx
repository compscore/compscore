export type check = {
  name: string;
  edges: {
    status?: status;
    credential?: credential;
  };
};

export type status = {
  error: string;
  status: string;
  time: string;
  edges: {
    check?: check;
    team?: team;
    round?: round;
  };
};

export type credential = {
  password: string;
  edges: {
    check?: check;
    team?: team;
  };
};

export type team = {
  number: number;
  name: string;
  edges: {
    status?: status;
    credential?: credential;
    round?: round;
  };
};

export type round = {
  number: number;
  edges: {
    status?: status;
    team?: team;
  };
};
