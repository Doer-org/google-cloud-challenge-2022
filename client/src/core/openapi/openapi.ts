/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/events": {
    /**
     * Create a new Event 
     * @description Creates a new Event and persists it to storage.
     */
    post: operations["createEvent"];
  };
  "/events/{id}": {
    /**
     * Find a Event by ID 
     * @description Finds the Event with the requested ID and returns it.
     */
    get: operations["readEvent"];
    /**
     * Deletes a Event by ID 
     * @description Deletes the Event with the requested ID.
     */
    delete: operations["deleteEvent"];
    /**
     * Updates a Event 
     * @description Updates a Event and persists changes to storage.
     */
    patch: operations["updateEvent"];
  };
  "/events/{id}/admin": {
    /**
     * Find the attached User 
     * @description Find the attached User of the Event with the given ID
     */
    get: operations["readEventAdmin"];
  };
  "/events/{id}/comments": {
    /**
     * List attached Comments 
     * @description List attached Comments.
     */
    get: operations["listEventComments"];
  };
  "/events/{id}/participants": {
    post: operations["postEventParticipants"];
    parameters: {
      path: {
        id: number;
      };
    };
  };
  "/events/{id}/state": {
    patch: operations["patchState"];
    parameters: {
      path: {
        id: number;
      };
    };
  };
  "/events/{id}/users": {
    /**
     * List attached Users 
     * @description List attached Users.
     */
    get: operations["listEventUsers"];
  };
  "/users": {
    /**
     * Create a new User 
     * @description Creates a new User and persists it to storage.
     */
    post: operations["createUser"];
  };
  "/users/{id}": {
    /**
     * Find a User by ID 
     * @description Finds the User with the requested ID and returns it.
     */
    get: operations["readUser"];
    /**
     * Deletes a User by ID 
     * @description Deletes the User with the requested ID.
     */
    delete: operations["deleteUser"];
    /**
     * Updates a User 
     * @description Updates a User and persists changes to storage.
     */
    patch: operations["updateUser"];
  };
  "/users/{id}/events": {
    /**
     * List attached Events 
     * @description List attached Events.
     */
    get: operations["listUserEvents"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    Comment: {
      /** Format: uuid */
      id: string;
      body: string;
      event: components["schemas"]["Event"];
      user: components["schemas"]["User"];
    };
    Event: {
      /** Format: uuid */
      id: string;
      name: string;
      detail?: string;
      location?: string;
      type: string;
      state: string;
      admin?: components["schemas"]["User"];
      users?: (components["schemas"]["User"])[];
      comments?: (components["schemas"]["Comment"])[];
    };
    EventCreate: {
      /** Format: uuid */
      id: string;
      name: string;
      detail?: string;
      location?: string;
      type: string;
      state: string;
    };
    EventRead: {
      /** Format: uuid */
      id: string;
      name: string;
      detail?: string;
      location?: string;
      type: string;
      state: string;
    };
    EventUpdate: {
      /** Format: uuid */
      id: string;
      name: string;
      detail?: string;
      location?: string;
      type: string;
      state: string;
    };
    Event_AdminRead: {
      /** Format: uuid */
      id: string;
      name: string;
      authenticated: boolean;
      mail?: string;
      icon?: string;
    };
    Event_CommentsList: {
      /** Format: uuid */
      id: string;
      body: string;
    };
    Event_UsersList: {
      /** Format: uuid */
      id: string;
      name: string;
      authenticated: boolean;
      mail?: string;
      icon?: string;
    };
    User: {
      /** Format: uuid */
      id: string;
      name: string;
      authenticated: boolean;
      mail?: string;
      icon?: string;
      events?: (components["schemas"]["Event"])[];
      comments?: (components["schemas"]["Comment"])[];
    };
    UserCreate: {
      /** Format: uuid */
      id: string;
      name: string;
      authenticated: boolean;
      mail?: string;
      icon?: string;
    };
    UserRead: {
      /** Format: uuid */
      id: string;
      name: string;
      authenticated: boolean;
      mail?: string;
      icon?: string;
    };
    UserUpdate: {
      /** Format: uuid */
      id: string;
      name: string;
      authenticated: boolean;
      mail?: string;
      icon?: string;
    };
    User_EventsList: {
      /** Format: uuid */
      id: string;
      name: string;
      detail?: string;
      location?: string;
      type: string;
      state: string;
    };
  };
  responses: {
    /** @description invalid input, data invalid */
    400: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: Record<string, never>;
        };
      };
    };
    /** @description insufficient permissions */
    403: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: Record<string, never>;
        };
      };
    };
    /** @description resource not found */
    404: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: Record<string, never>;
        };
      };
    };
    /** @description conflicting resources */
    409: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: Record<string, never>;
        };
      };
    };
    /** @description unexpected error */
    500: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: Record<string, never>;
        };
      };
    };
  };
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type external = Record<string, never>;

export interface operations {

  createEvent: {
    /**
     * Create a new Event 
     * @description Creates a new Event and persists it to storage.
     */
    /** @description Event to create */
    requestBody: {
      content: {
        "application/json": {
          name: string;
          detail?: string;
          location?: string;
          type: string;
          state: string;
          /** Format: uuid */
          admin?: string;
          users?: (string)[];
          comments?: (string)[];
        };
      };
    };
    responses: {
      /** @description Event created */
      200: {
        content: {
          "application/json": components["schemas"]["EventCreate"];
        };
      };
      400: components["responses"]["400"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  readEvent: {
    /**
     * Find a Event by ID 
     * @description Finds the Event with the requested ID and returns it.
     */
    parameters: {
        /** @description ID of the Event */
      path: {
        id: string;
      };
    };
    responses: {
      /** @description Event with requested ID was found */
      200: {
        content: {
          "application/json": components["schemas"]["EventRead"];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  deleteEvent: {
    /**
     * Deletes a Event by ID 
     * @description Deletes the Event with the requested ID.
     */
    parameters: {
        /** @description ID of the Event */
      path: {
        id: string;
      };
    };
    responses: {
      /** @description Event with requested ID was deleted */
      204: never;
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  updateEvent: {
    /**
     * Updates a Event 
     * @description Updates a Event and persists changes to storage.
     */
    parameters: {
        /** @description ID of the Event */
      path: {
        id: string;
      };
    };
    /** @description Event properties to update */
    requestBody: {
      content: {
        "application/json": {
          name?: string;
          detail?: string;
          location?: string;
          type?: string;
          state?: string;
          /** Format: uuid */
          admin?: string;
          users?: (string)[];
          comments?: (string)[];
        };
      };
    };
    responses: {
      /** @description Event updated */
      200: {
        content: {
          "application/json": components["schemas"]["EventUpdate"];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  readEventAdmin: {
    /**
     * Find the attached User 
     * @description Find the attached User of the Event with the given ID
     */
    parameters: {
        /** @description ID of the Event */
      path: {
        id: string;
      };
    };
    responses: {
      /** @description User attached to Event with requested ID was found */
      200: {
        content: {
          "application/json": components["schemas"]["Event_AdminRead"];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  listEventComments: {
    /**
     * List attached Comments 
     * @description List attached Comments.
     */
    parameters: {
        /** @description what page to render */
        /** @description item count to render per page */
      query?: {
        page?: number;
        itemsPerPage?: number;
      };
        /** @description ID of the Event */
      path: {
        id: string;
      };
    };
    responses: {
      /** @description result Events list */
      200: {
        content: {
          "application/json": (components["schemas"]["Event_CommentsList"])[];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  postEventParticipants: {
    responses: {
      200: never;
    };
  };
  patchState: {
    responses: {
      200: never;
    };
  };
  listEventUsers: {
    /**
     * List attached Users 
     * @description List attached Users.
     */
    parameters: {
        /** @description what page to render */
        /** @description item count to render per page */
      query?: {
        page?: number;
        itemsPerPage?: number;
      };
        /** @description ID of the Event */
      path: {
        id: string;
      };
    };
    responses: {
      /** @description result Events list */
      200: {
        content: {
          "application/json": (components["schemas"]["Event_UsersList"])[];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  createUser: {
    /**
     * Create a new User 
     * @description Creates a new User and persists it to storage.
     */
    /** @description User to create */
    requestBody: {
      content: {
        "application/json": {
          name: string;
          authenticated: boolean;
          mail?: string;
          icon?: string;
          events?: (string)[];
          comments?: (string)[];
        };
      };
    };
    responses: {
      /** @description User created */
      200: {
        content: {
          "application/json": components["schemas"]["UserCreate"];
        };
      };
      400: components["responses"]["400"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  readUser: {
    /**
     * Find a User by ID 
     * @description Finds the User with the requested ID and returns it.
     */
    parameters: {
        /** @description ID of the User */
      path: {
        id: string;
      };
    };
    responses: {
      /** @description User with requested ID was found */
      200: {
        content: {
          "application/json": components["schemas"]["UserRead"];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  deleteUser: {
    /**
     * Deletes a User by ID 
     * @description Deletes the User with the requested ID.
     */
    parameters: {
        /** @description ID of the User */
      path: {
        id: string;
      };
    };
    responses: {
      /** @description User with requested ID was deleted */
      204: never;
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  updateUser: {
    /**
     * Updates a User 
     * @description Updates a User and persists changes to storage.
     */
    parameters: {
        /** @description ID of the User */
      path: {
        id: string;
      };
    };
    /** @description User properties to update */
    requestBody: {
      content: {
        "application/json": {
          name?: string;
          authenticated?: boolean;
          mail?: string;
          icon?: string;
          events?: (string)[];
          comments?: (string)[];
        };
      };
    };
    responses: {
      /** @description User updated */
      200: {
        content: {
          "application/json": components["schemas"]["UserUpdate"];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  listUserEvents: {
    /**
     * List attached Events 
     * @description List attached Events.
     */
    parameters: {
        /** @description what page to render */
        /** @description item count to render per page */
      query?: {
        page?: number;
        itemsPerPage?: number;
      };
        /** @description ID of the User */
      path: {
        id: string;
      };
    };
    responses: {
      /** @description result Users list */
      200: {
        content: {
          "application/json": (components["schemas"]["User_EventsList"])[];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
}