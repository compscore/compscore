export type Check = {
  name: string;
  teams: [number];
};

export type Scoreboard = {
  round: number;
  checks: [Check];
};
