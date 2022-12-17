<?php

namespace App\Models;

use App\Enums\PaymentTypes;
use App\Models\Concerns\HasUuid;
use App\Payment\PaymentType;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Employee extends Model
{
    use HasFactory;
    use HasUuid;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'uuid',
        'full_name',
        'email',
        'department_id',
        'job_title',
        'payment_type',
        'salary',
        'hourly_rate',
    ];

    /**
     * The attributes that should be cast to native types.
     *
     * @var array
     */
    protected $casts = [
        'id' => 'integer',
        'department_id' => 'integer',
    ];

    public function paychecks()
    {
        return $this->hasMany(Paycheck::class);
    }

    public function timelogs()
    {
        return $this->hasMany(Timelog::class);
    }

    public function department()
    {
        return $this->belongsTo(Department::class);
    }

    public function getRouteKeyName(): string
    {
        return 'uuid';
    }

    public function getPaymentTypeAttribute(): PaymentType
    {
        return PaymentTypes::from($this->original['payment_type'])->makePaymentType($this);
    }
}
