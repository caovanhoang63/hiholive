export interface Stream {
    id: number,
    title: string,
    channelId: number,
    state: string
    status: number,
}

export const StatePending = "pending"
export const StateScheduled = "scheduled"
export const StateRunning = "running"
export const StateEnded = "ended"
