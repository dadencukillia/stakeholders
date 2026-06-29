export type ProjectStatus = "working" | "finished" | "paused" | "dropped";

export interface ProjectCardProps {
  id: string,
  shareUrl: string,
  title: string,
  description: string,
  commits: number,
  branches: number,
  workingHours: number,
  manualStatus: string,
  status: ProjectStatus,
  codeAccess: boolean,
} 
