<?php

namespace App\Http\Resources;

use Illuminate\Support\Arr;
use TiMacDonald\JsonApi\JsonApiResource;
use TiMacDonald\JsonApi\Link;

class PostResource extends JsonApiResource
{
    public function toAttributes($request): array
    {
        return Arr::except($this->resource->toArray(), ['id', 'author', 'comments']);
    }

    public function toRelationships($request): array
    {
        return [
            'author' => fn () => new UserResource($this->author),
            'comments' => fn () => CommentResource::collection($this->comments),
        ];
    }

    public function toLinks($request): array
    {
        return [
            Link::self(route('posts.show', $this->resource)),
        ];
    }
}
