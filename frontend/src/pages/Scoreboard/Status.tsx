import { Box } from "@mui/material";
import { useParams } from "react-router";
import StatusScoreboard from "../../components/Scoreboard/Status";

export default function StatusScoreboardPage() {
  const { check: check, team: team } = useParams() as unknown as {
    check: string;
    team: number;
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
      <StatusScoreboard check={check} team={team} />
    </Box>
  );
}
