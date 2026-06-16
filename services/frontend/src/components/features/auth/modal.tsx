import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import type React from "react";

export const ModalAuthForm = ({
  action,
  children,
  formSlot,
}: {
  action: "login"|"register",
  children: React.ReactNode,
  formSlot: React.ReactNode
}) => {
  return (<Dialog>
    <DialogTrigger asChild><div>{ children }</div></DialogTrigger>
    <DialogContent>
      <DialogHeader>
      <DialogTitle>{ action.toLowerCase().replace(/\b\w/g, l => l.toUpperCase()) }</DialogTitle>
      </DialogHeader>
      { formSlot }
    </DialogContent>
  </Dialog>);
};
