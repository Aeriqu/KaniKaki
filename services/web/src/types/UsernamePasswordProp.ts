import { Dispatch, SetStateAction } from "react"

export interface UsernamePasswordProp {
  username: string
  password: string
  setErrorState: Dispatch<SetStateAction<string>>
}