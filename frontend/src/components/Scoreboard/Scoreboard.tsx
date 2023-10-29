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
import { enqueueSnackbar } from "notistack";
import { useEffect, useState } from "react";
import { api_url, short_refresh } from "../../config";
import { Scoreboard } from "../../models/Scoreboard/Scoreboard";

export default function ScoreBoard() {
  const [data, setData] = useState<Scoreboard>();

  useEffect(() => {
    const fetchData = async () => {
      fetch(`${api_url}/api/scoreboard`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then(async (res) => {
          if (res.status === 200) {
            setData((await res.json()) as Scoreboard);
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

  const [highlightedTeam, setHighlightedTeam] = useState<number | null>(null);
  const [highlightedCheck, setHighlightedCheck] = useState<string | null>(null);

  const getBackgroundColor = (status: number, team: number, check: string) => {
    if (highlightedCheck === null && highlightedTeam === null) {
      if (status === 0) {
        return "#f44336";
      } else if (status === 1) {
        return "#4caf50";
      }
      return "#999891";
    }
    if (highlightedCheck === check || highlightedTeam === team) {
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
      >
        Scoreboard
      </Typography>
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
        }}
      >
        {data && data.round > 10 && (
          <DoubleArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = "/scoreboard/round/" + (data.round - 10);
            }}
          />
        )}
        {data && data.round > 1 && (
          <ArrowLeftIcon
            sx={{
              cursor: "pointer",
            }}
            onClick={() => {
              window.location.href = "/scoreboard/round/" + (data.round - 1);
            }}
          />
        )}
        <Typography component='h1' variant='h5'>
          Round {data?.round}
        </Typography>
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
              <TableCell size='small'>
                <Typography variant='subtitle2'>Team</Typography>
              </TableCell>
              {data?.checks[0].status.map((_, team) => (
                <TableCell
                  key={team + 1}
                  align='center'
                  size='small'
                  onMouseEnter={() => {
                    setHighlightedTeam(team + 1);
                  }}
                  onMouseLeave={() => {
                    setHighlightedTeam(null);
                  }}
                  sx={{
                    backgroundColor:
                      highlightedTeam === team + 1 ? "#343434" : "transparent",
                  }}
                  onClick={() => {
                    window.location.href = "/scoreboard/team/" + (team + 1);
                  }}
                >
                  <Typography variant='subtitle2'>{team + 1}</Typography>
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.checks.map((check) => (
              <TableRow>
                <TableCell
                  key={check.name}
                  size='small'
                  onMouseEnter={() => {
                    setHighlightedCheck(check.name);
                  }}
                  onMouseLeave={() => {
                    setHighlightedCheck(null);
                  }}
                  sx={{
                    backgroundColor:
                      highlightedCheck === check.name
                        ? "#343434"
                        : "transparent",
                  }}
                  onClick={() => {
                    window.location.href = `/scoreboard/check/${check.name}`;
                  }}
                >
                  {check.name}
                </TableCell>
                {check.status.map((status, team) => (
                  <TableCell
                    key={team + "-" + check.name}
                    size='small'
                    sx={{
                      backgroundColor: getBackgroundColor(
                        status,
                        team + 1,
                        check.name
                      ),
                    }}
                    onMouseEnter={() => {
                      setHighlightedTeam(team + 1);
                      setHighlightedCheck(check.name);
                    }}
                    onMouseLeave={() => {
                      setHighlightedTeam(null);
                      setHighlightedCheck(null);
                    }}
                    onClick={() => {
                      window.location.href = `/scoreboard/status/${team + 1}/${
                        check.name
                      }`;
                    }}
                  ></TableCell>
                ))}
              </TableRow>
            ))}
            <TableRow>
              <TableCell size='small'>
                <Typography variant='subtitle2'>Score</Typography>
              </TableCell>
              {data?.scores.map((score, team) => (
                <TableCell
                  key={"score" + team + 1}
                  size='small'
                  align='center'
                  onMouseEnter={() => {
                    setHighlightedTeam(team + 1);
                  }}
                  onMouseLeave={() => {
                    setHighlightedTeam(null);
                  }}
                  sx={{
                    backgroundColor:
                      highlightedTeam === team + 1 ? "#343434" : "transparent",
                  }}
                >
                  <Typography variant='subtitle2'>{score}</Typography>
                </TableCell>
              ))}
            </TableRow>
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
