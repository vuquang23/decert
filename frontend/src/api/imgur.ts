const upload = async (image: File) => {
  const headers = new Headers();
  headers.append("Authorization", "Client-ID 1744f29d581f1eb");

  const formData = new FormData();
  formData.append("image", image);

  const response = await fetch("https://api.imgur.com/3/image", {
    method: "POST",
    headers: headers,
    body: formData,
    redirect: "follow",
  });

  const json = await response.json();
  if (!response.ok || !("data" in json)) {
    throw Error(json.message ?? response.statusText);
  }
  return json.data.link as string;
};

export { upload };
