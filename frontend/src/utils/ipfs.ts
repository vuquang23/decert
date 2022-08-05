import axios from "axios";
import FormData from "form-data";

export async function pushToIpfs(readStream: any): Promise<string | Error> {
  const JWT =
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiJkOTliNWZjNS0xMjJhLTQwMTUtYWMxNy1iMWZhMjdmOGNiNzMiLCJlbWFpbCI6InRlc3Rud2JsY0BnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwicGluX3BvbGljeSI6eyJyZWdpb25zIjpbeyJpZCI6IkZSQTEiLCJkZXNpcmVkUmVwbGljYXRpb25Db3VudCI6MX0seyJpZCI6Ik5ZQzEiLCJkZXNpcmVkUmVwbGljYXRpb25Db3VudCI6MX1dLCJ2ZXJzaW9uIjoxfSwibWZhX2VuYWJsZWQiOmZhbHNlLCJzdGF0dXMiOiJBQ1RJVkUifSwiYXV0aGVudGljYXRpb25UeXBlIjoic2NvcGVkS2V5Iiwic2NvcGVkS2V5S2V5IjoiODM3MzVlYWNkMmMyMGU1Y2MwNTUiLCJzY29wZWRLZXlTZWNyZXQiOiJmOWM5NzI2ZDM2YmJkMjhkMjYyMzM1YWE3YmE5NDY5NDk2MjBjN2MyMTk5MGY2ODE0OGQ5OTljYzQ5ODcwODQzIiwiaWF0IjoxNjU5NjgxMDg3fQ.JgYZ8_HbwD7VY4Q2FxDc6NavQog2wdu30j-VZLDdpVU";

  const data = new FormData();
  data.append("file", readStream);
  data.append("pinataOptions", '{"cidVersion": 1}');

  const config = {
    method: "post",
    url: "https://api.pinata.cloud/pinning/pinFileToIPFS",
    headers: {
      Authorization: `Bearer ${JWT}`,
      ...data.getHeaders(),
    },
    data: data,
  };
  try {
    const response = await axios(config);
    return `https://gateway.pinata.cloud/ipfs/${response.data["IpfsHash"]}`;
  } catch (e) {
    console.log(e);
    return new Error("Error to push to ipfs");
  }
}
