import { Injectable } from '@angular/core';
import { Folder } from '../models/folder';
import { File } from "../models/file";
import { User } from "../models/user";

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor() { }

  createUser(user: User){

  }
  login(login: Login){}
  createFolder(folder: Folder){

  }
  createFile(file: File){

  }
  listFolders(userId: string): Folder[]{
     let folders: Folder[]=[];

     return folders;
  }
}
