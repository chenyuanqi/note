<?php
namespace App\Http\Resources;

use App\VOs\Amount;
use Illuminate\Http\Request;
use TiMacDonald\JsonApi\JsonApiResource;

class EmployeeResource extends JsonApiResource
{
    /**
     * Transform the resource into an array.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array|\Illuminate\Contracts\Support\Arrayable|\JsonSerializable
     */
    public function toAttributes($request): array
    {
        return [
            'name' => $this->full_name,
            'email' => $this->email,
            'jobTitle' => $this->job_title,
            'payment' => [
                'type' => $this->payment_type->type(),
                'amount' => Amount::from($this->payment_type->amount())->toArray(),
            ],
        ];
    }

    public function toId(Request $request): string
    {
       return $this->uuid; 
    }
}
