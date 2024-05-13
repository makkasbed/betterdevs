import { Injectable } from '@angular/core';
import { Folder } from '../models/folder';
import { File } from "../models/file";
import { User } from "../models/user";
import { HttpClient } from '@angular/common/http';
import { Response } from '../models/response';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  url: string = "http://localhost:18085";

  constructor(private http: HttpClient) { }

  createUser(user: User){
    return this.http.post<Response>(this.url+"/users", user);
  }
  login(login: Login){
    return this.http.post<Response>(this.url+"/login", login);
  }
  createFolder(folder: Folder){
    return this.http.post<Response>(this.url+"/folders", folder);
  }
  createFile(file: File){
    return this.http.post<Response>(this.url+"/files", file);
  }
  listFolders(userId: string){
     return this.http.get<Folder[]>(this.url+`/folders/${userId}`)
  }
}
