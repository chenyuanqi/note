<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateEmployeesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::disableForeignKeyConstraints();

        Schema::create('employees', function (Blueprint $table) {
            $table->id();
            $table->uuid('uuid');
            $table->string('full_name', 100);
            $table->string('email', 100)->index();
            $table->foreignId('department_id')->constrained();
            $table->string('job_title', 50);
            $table->string('payment_type', 20);
            $table->unsignedInteger('salary')->nullable();
            $table->unsignedInteger('hourly_rate')->nullable();
            $table->timestamps();
        });

        Schema::enableForeignKeyConstraints();
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('employees');
    }
}
