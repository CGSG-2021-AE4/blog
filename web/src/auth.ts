
export function Login( username: string, token: string ) {
  localStorage.setItem("ae4blog-authorized", "true")
  localStorage.setItem("ae4blog-username", username)
  localStorage.setItem("ae4blog-token", token)
}

export function Logout() {
  localStorage.setItem("ae4blog-authorized", "false")
  localStorage.setItem("ae4blog-username", "")
  localStorage.setItem("ae4blog-token", "")
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

export function GetToken(): string {
  let u = localStorage.getItem("ae4blog-token")
  if (u == null) {
    return ""
  }
  return u
}