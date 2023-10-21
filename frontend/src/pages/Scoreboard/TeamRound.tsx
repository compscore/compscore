import { Box } from "@mui/material";
import { useParams } from "react-router-dom";
import TeamRoundScoreboard from "../../components/Scoreboard/TeamRound";

export default function TeamRoundScoreboardPage() {
  const { team: team, round: round } = useParams() as unknown as {
    team: string;
    round: string;
  };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <TeamRoundScoreboard team={team} round={round} />
    </Box>
  );
}
