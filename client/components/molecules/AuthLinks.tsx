import { m } from 'vitest/dist/index-2d10c3fd';
import { createUser } from '../../core/api/user/create';
import { Button } from '../atoms/text/Button';
import { LinkTo } from '../atoms/text/LinkTo';
type TProps = {
  auth: boolean;
  changeAuth: (auth: boolean) => void;
};
export const AuthLinks = ({ auth, changeAuth }: TProps) => {
  // TODO:ログインhooksを使ってonClickの中へ
  // TODO:サインインhooksを使ってonClickの中へ
  // TODO:ログアウトhooksを使ってonClickの中へ
  // TODO: User ID 
  const login = () => {
    createUser(
      (ok) => {},
      (err) => {}
    )
  }
  return (
    <div className="my-10 grid grid-cols-1 gap-2">
      {!auth ? (
        <>
          <Button  
            className="m-1" 
            onClick={() => {
              
              changeAuth(true)
            }}
            >
            ログイン
          </Button>
          <Button onClick={() => {}} className="m-1">
            サインイン
          </Button>
        </>
      ) : (
        <>
          <LinkTo href="/event/new" className="m-1">
            イベント募集ページへ
          </LinkTo>
          <LinkTo href="/event" className="m-1">
            イベント一覧ページへ
          </LinkTo>
          <Button onClick={() => changeAuth(false)} className="m-1">
            ログアウト
          </Button>
        </>
      )}
    </div>
  );
};
