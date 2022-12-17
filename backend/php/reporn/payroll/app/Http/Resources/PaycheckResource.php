<?php
namespace App\Http\Resources;

use App\VOs\Amount;
use Illuminate\Http\Request;
use Illuminate\Support\Arr;
use TiMacDonald\JsonApi\JsonApiResource;

class PaycheckResource extends JsonApiResource
{
    public function toAttributes($request): array
    {
        return [
            'amount' => Amount::from($this->net_amount)->toArray(),
            'payed_at' => $this->payed_at,
        ];
    }

    public function toId(Request $request): string
    {
        return $this->uuid;
    }

    public function toRelationships(Request $request): array
    {
        return [
            'employee' => fn() => EmployeeResource::make($this->employee)
        ];
    }
}