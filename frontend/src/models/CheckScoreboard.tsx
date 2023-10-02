export type TeamScoreboard = {
  round: number;
  teams: [
    {
      name: string;
      status: [number];
    }
  ];
};
