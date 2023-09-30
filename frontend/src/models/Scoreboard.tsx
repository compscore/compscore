export type Scoreboard = {
  round: number;
  scores: [number];
  checks: [
    {
      name: string;
      teams: [number];
    }
  ];
};
