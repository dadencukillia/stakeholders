export type ProjectStatus = "working" | "finished" | "paused" | "dropped";

export interface ProjectCardProps {
  id: string,
  shareUrl: string,
  title: string,
  description: string,
  commits: number,
  branches: number,
  wakaHours: number,
  status: ProjectStatus,
  codeAccess: boolean,
} 
