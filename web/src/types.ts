
export class articleHeader {
  id: string;
	title: string;
  authorId: string;
  authorUsername: string;
  created: string;
  contentId: string;
}

// Request and responses
export class errorResp {
  err: string;
}

export class msgResp {
  msgResp: string;
}

export class editArticleReq {
	id: string;
	title: string;
	content: string;
}
