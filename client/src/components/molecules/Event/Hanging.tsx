import React from 'react';

export const Hanging = () => {
  return (
    <div className="flex justify-center items-center flex-col -mb-8">
      <div className="w-8 h-8 rounded-full bg-slate-400 m-4 z-20"></div>
      <div className="flex justify-center">
        <div className="z-10">
          <div className="border border-black w-24 -rotate-45"></div>
          <div className="w-3 h-3 rounded-full bg-slate-400 mr-auto ml-2 mt-6 z-20 relative"></div>
        </div>
        <div className="z-10">
          <div className="border border-black w-24 rotate-45"></div>
          <div className="w-3 h-3 rounded-full bg-slate-400 ml-auto mr-2 mt-6 z-20 relative"></div>
        </div>
      </div>
    </div>
  );
};
