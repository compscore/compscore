import { Box, Typography } from "@mui/material";
import CheckScoreboard from "../components/CheckScoreboard";
import { useParams } from "react-router-dom";

export default function CheckScoreBoard() {
  const { check: check } = useParams() as { check: string };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Typography
        component='h1'
        variant='h3'
        fontWeight={700}
        sx={{
          marginTop: 5,
        }}
        onClick={() => {
          window.location.href = "/scoreboard";
        }}
      >
        {check}
      </Typography>
      <CheckScoreboard check={check} />
    </Box>
  );
}
