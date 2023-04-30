import { FCProps } from "@/types/FCProps";
import Header from "../modules/header/Header";

export default function HeaderLayout({ children }: FCProps) {
  return (
    <div>
      <Header />
      <main>{children}</main>
    </div>
  )
}