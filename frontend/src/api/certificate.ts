interface Certificate {
  id: number;
  title: string;
  receiver: string;
  description: string;
  issuedAt: Date;
  expiredAt: Date;
  imgUrl: string;
  revokeAt?: Date;
  revokeReason?: String;
}

const today = new Date(Date.now());

const mockData: Certificate[] = Array.from(Array(30).keys(), (_, index) => ({
  id: Math.random() * 100,
  title: `Sinh viên ${index} tốt`,
  receiver: "",
  description: "Đã đạt thành tích xuất sắc trong học tập",
  issuedAt: today,
  expiredAt: new Date(
    new Date(today).setDate(today.getDate() + Math.floor(Math.random() * 3))
  ),
  imgUrl: "https://picsum.photos/1000/1000",
}));

const isExpired = (cert: Certificate) =>
  cert.expiredAt.getTime() - Date.now() <= 0;

const readAll = (receiver: string) => Promise.resolve(mockData);

export type { Certificate };
export { readAll, isExpired };
