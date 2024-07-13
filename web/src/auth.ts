
export function Login( id: string, username: string, token: string ) {
  localStorage.setItem("ae4blog-authorized", "true")
  localStorage.setItem("ae4blog-username", username)
  localStorage.setItem("ae4blog-id", id)
  localStorage.setItem("ae4blog-token", token)
}

export function Logout() {
  localStorage.setItem("ae4blog-authorized", "false")
  localStorage.setItem("ae4blog-username", "")
  localStorage.setItem("ae4blog-token", "")
  localStorage.setItem("ae4blog-id", "")
}

export function IsAuthorized(): boolean {
  return (localStorage.getItem("ae4blog-authorized") == "true")
}

export function GetUsername(): string { // Maybe shorten later
  let u = localStorage.getItem("ae4blog-username")
  if (u == null) {
    return ""
  }
  return u
}

export function GetId(): string { // Maybe shorten later
  let id = localStorage.getItem("ae4blog-id")
  if (id == null) {
    return ""
  }
  return id
}

export function GetToken(): string {
  let u = localStorage.getItem("ae4blog-token")
  if (u == null) {
    return ""
  }
  return u
}