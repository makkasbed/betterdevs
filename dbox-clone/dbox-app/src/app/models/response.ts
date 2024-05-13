import { Folder } from "./folder";

export interface Response{
  status: number;
  message: String;
  id: String;
  folders: Folder[]
}
