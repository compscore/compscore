import { Box } from "@mui/material";
import { useParams } from "react-router-dom";
import CheckScoreboard from "../../components/CheckScoreboard";

export default function CheckScoreboardPage() {
  const { check: check } = useParams() as { check: string };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <CheckScoreboard check={check} />
    </Box>
  );
}
