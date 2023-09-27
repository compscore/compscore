export class JWT {
  Team: string;
  exp: number;
  constructor(team: string, exp: number) {
    this.Team = team;
    this.exp = exp;
  }
}
