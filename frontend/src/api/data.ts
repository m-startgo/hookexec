import { MyFetch } from './http';

export function GetData(data: unknown): Promise<unknown> {
  return MyFetch({
    Path: '/api/data',
    Data: data,
    Method: 'GET',
  });
}

export function PostData(data: unknown): Promise<unknown> {
  return MyFetch({
    Path: '/api/data',
    Data: data,
    Method: 'POST',
  });
}
