export type Scoreboard = {
  round: number;
  scores: [number];
  checks: [
    {
      name: string;
      status: [number];
    }
  ];
};
