import ArrowLeftIcon from "@mui/icons-material/KeyboardArrowLeft";
import DoubleArrowLeftIcon from "@mui/icons-material/KeyboardDoubleArrowLeft";
import {
  Box,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import { useEffect, useState } from "react";
import { StatusScoreboard } from "../../models/Scoreboard/Status";
import { api_url, short_refresh } from "../../config";
import { enqueueSnackbar } from "notistack";

type props = {
  check: string;
  team: number;
};

export default function StatusScoreboardComponent({ check, team }: props) {
  const [data, setData] = useState<StatusScoreboard>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`${api_url}/api/scoreboard/status/${team}/${check}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          if (res.status === 200) {
            setData((await res.json()) as StatusScoreboard);
          } else {
            enqueueSnackbar("Encountered an error", { variant: "error" });
          }
        })
        .catch((err) => {
          enqueueSnackbar("Encountered an error: " + err, { variant: "error" });
          console.log(err);
        });
    };

    fetchData();

    const pollingInterval = setInterval(fetchData, short_refresh);

    return () => clearInterval(pollingInterval);
  }, []);

  const [highlightedRound, setHighlightedRound] = useState<number | null>(null);

  const getBackgroundColor = (status: number, round: number) => {
    if (highlightedRound === null) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (highlightedRound === round) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (status === 0) {
      return "#f26359";
    } else if (status === 1) {
      return "#75ad77";
    }
    return "#cecdc6";
  };

  return (
    <>
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
        Team {team} - {check}
      </Typography>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
        }}
      >
        {data && data[0].round >= 10 && (
          <DoubleArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/status/${team}/${check}/${
                data[0].round - 10
              }`;
            }}
          />
        )}
        {data && data[0].round >= 1 && (
          <ArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = `/scoreboard/status/${team}/${check}/${
                data[0].round - 1
              }`;
            }}
          />
        )}
        {data && (
          <Typography component='h1' variant='h5'>
            Round {data[0].round}
          </Typography>
        )}
        <ArrowLeftIcon
          sx={{
            visibility: "hidden",
          }}
        />
        <ArrowLeftIcon
          sx={{
            visibility: "hidden",
          }}
        />
      </Box>
      <Box m={2}></Box>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              {data?.map((status) => (
                <TableCell
                  size='small'
                  align='center'
                  key={"round-" + status.round}
                >
                  {status.round}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            <TableRow>
              {data?.map((status) => (
                <TableCell
                  size='medium'
                  align='center'
                  key={"status-" + status.round}
                  sx={{
                    backgroundColor: getBackgroundColor(
                      status.status,
                      status.round
                    ),
                  }}
                  onMouseEnter={() => {
                    setHighlightedRound(status.round);
                  }}
                  onMouseLeave={() => {
                    setHighlightedRound(null);
                  }}
                >
                  <Box height={75}></Box>
                </TableCell>
              ))}
            </TableRow>
          </TableBody>
        </Table>
      </TableContainer>
      <Box m={2}></Box>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell size='small' align='center'>
                <Typography variant='subtitle2'>Round</Typography>
              </TableCell>
              <TableCell size='small' align='center'>
                <Typography variant='subtitle2'>Time</Typography>
              </TableCell>
              <TableCell size='small' align='center'>
                <Typography variant='subtitle2'>Status</Typography>
              </TableCell>
              <TableCell size='small' align='center'>
                <Typography variant='subtitle2'>Error</Typography>
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.map((status) => (
              <TableRow
                onMouseEnter={() => {
                  setHighlightedRound(status.round);
                }}
                onMouseLeave={() => {
                  setHighlightedRound(null);
                }}
                sx={{
                  backgroundColor:
                    highlightedRound === status.round
                      ? "#343434"
                      : "transparent",
                }}
              >
                <TableCell size='small' align='center'>
                  {status.round}
                </TableCell>
                <TableCell size='small' align='center'>
                  {status.time}
                </TableCell>
                <TableCell size='small' align='center'>
                  {status.status === 0
                    ? "Down"
                    : status.status === 1
                    ? "Up"
                    : "Unknown"}
                </TableCell>
                <TableCell size='small' align='center'>
                  {status.error}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
