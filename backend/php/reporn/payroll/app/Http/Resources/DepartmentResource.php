<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use TiMacDonald\JsonApi\JsonApiResource;

class DepartmentResource extends JsonApiResource
{
    public function toAttributes($request): array
    {
        return [
            'name' => $this->name,
            'description' => $this->description,
        ];
    }

    public function toRelationships(Request $request): array
    {
        return [
            'employees' => fn() => EmployeeResource::collection($this->employees)
        ];
    }

    public function toLinks($request): array
    {
        return [
            'self' => route('departments.show', ['department' => $this->uuid]),
        ];
    }
}
