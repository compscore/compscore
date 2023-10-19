import { Box } from "@mui/material";
import { useParams } from "react-router";
import StatusRoundScoreboard from "../../components/Scoreboard/StatusRound";

export default function StatusRoundScoreboardPage() {
  const {
    check: check,
    team: team,
    round: round,
  } = useParams() as unknown as {
    check: string;
    team: string;
    round: string;
  };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        width: "80%",
        margin: "auto",
      }}
    >
      <StatusRoundScoreboard check={check} team={team} round={round} />
    </Box>
  );
}
