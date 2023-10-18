import { Box } from "@mui/material";
import Scoreboard from "../../components/Scoreboard/Scoreboard";

export default function ScoreboardPage() {
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Scoreboard />
      <Box m={2}></Box>
    </Box>
  );
}
