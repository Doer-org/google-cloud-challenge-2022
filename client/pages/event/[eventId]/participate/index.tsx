import { useState } from 'react';
import { useRouter } from 'next/router';
import { MyHead } from '../../../../components/templates/shared/MyHead';
import { BasicTemplate } from '../../../../components/templates/shared/BasicTemplate';
import { FormWrapper } from '../../../../components/atoms/form/FormWrapper';
import { Input } from '../../../../components/atoms/form/Input';
import { Button } from '../../../../components/atoms/text/Button';
import { EventInfo } from '../../../../components/molecules/EventInfo';
import { EventConfirmModal } from '../../../../components/molecules/Event/EventConfirmModal';
import Head from 'next/head';
import { getEventInfo, tryGetEventInfo } from '../../../../core/api/event/getInfo';
import { Event } from '../../../../core/types/event';
import { flow, pipe } from 'fp-ts/lib/function'
import * as TE from 'fp-ts/TaskEither'
import { joinEvent } from '../../../../core/api/event/join';

export async function getServerSideProps(context: any) {
  const eventId = context.query.eventId;
  //   const dummyEvent = {
  //     id: '1',
  //     eventName: 'ラーメン',
  //     detail: 'ターメン行きたいんや！！！！',
  //     position: 'lat:100/lng:200',
  //     capacity: 1,
  //   }; 
  //   return {
  //     props: {
  //       news: dummyEvent,
  //     },
  //   };
  // TODO: ここでeventIdを元にevent情報をとってきてpropsで返す
  // 型も付けといてもらえると助かる！！  
  return pipe(
    eventId,
    tryGetEventInfo,
    TE.match(
      (err) => { 
        throw err
      },
      (ok) => { 
        return {
          props: { 
            news: ok,
          },
        } 
      }
      )
  )() 
}

export default function Participate(props: any) {
  // ここは参加者がみるただの参加フォーム
  // TODO: SSRで実装してリンクを貼った時にOGPを表示させるようにする
  // TODO: event参加hooksをonClickへ

  const [isConfirm, setIsConfirm] = useState(false);
  const joinApi = joinEvent(
    () => { setIsConfirm(true) },
    () => {}
  )
  const [name, setName] = useState('');
  const [word, setWord] = useState('');
  const eventId = useRouter().query.eventId;
  console.log(props);
  return (
    <>
      <MyHead title="募集タイトルを入れる" description="" />
      <Head>
        <meta property="og:title" content={props.eventName} />
        <meta property="og:description" content={props.detail} />
        {/* <meta property="og:image" content={} /> */}
        <meta name="twitter:card" content="summary_large_image" />
      </Head>
      <BasicTemplate className="text-center">
        {isConfirm ? (
          <EventConfirmModal
            onParticipate={() => {}}
            onCancel={() => setIsConfirm(false)}
            eventId={Number(eventId as string)}
          />
        ) : (
          <>
            <EventInfo />
            <FormWrapper>
              <Input
                type="text"
                label="名前"
                content={name}
                changeContent={setName}
                required={true}
              />
              <Input
                type="text"
                label="ひとこと"
                content={word}
                changeContent={setWord}
                required={true}
              />
              <Button
                className="flex m-auto my-5"
                onClick={() => {
                  // setIsConfirm(true) 
                  joinApi({
                    event_id: eventId as string,
                    participant_name: name,
                    comment:word
                  })
                }}
              >
                参加する
              </Button>
            </FormWrapper>
          </>
        )}
      </BasicTemplate>
    </>
  );
}
