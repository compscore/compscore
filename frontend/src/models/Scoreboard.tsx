export type Status = {
  name: string;
  checks: [number];
};

export type Scoreboard = {
  round: number;
  checks: [Status];
};
