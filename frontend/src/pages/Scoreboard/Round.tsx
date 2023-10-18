import { Box } from "@mui/material";
import { useParams } from "react-router-dom";
import RoundScoreboard from "../../components/Scoreboard/Round";

export default function RoundScoreboardPage() {
  const { round: round } = useParams() as { round: string };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <RoundScoreboard round={round} />
    </Box>
  );
}
