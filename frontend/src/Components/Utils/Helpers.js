/* ------------------------------------ Helpers Utils --------------------------------------*/

// Format time.Duration Golang value sent from api back to layout: "2h35m"
function formatDuration(durationNanoseconds) {
  const duration = parseInt(durationNanoseconds, 10);
  const hours = Math.floor(duration / 3600000000000);
  const minutes = Math.floor((duration % 3600000000000) / 60000000000);

  return `${hours}h${minutes}m`;
}

export { formatDuration };
