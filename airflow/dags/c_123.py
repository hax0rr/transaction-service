
from airflow import DAG
from airflow.operators.dummy_operator import DummyOperator
from datetime import datetime, timedelta

default_args = {'owner': 'airflow', 'depends_on_past': False, 'retries': 1, 'retry_delay': datetime.timedelta(seconds=60)}

dag = DAG(
    dag_id='c_123',
    default_args=default_args,
    schedule_interval='None',
    start_date=2025-01-31 23:00:00,
    catchup=False
)

campaignId = 123

start = DummyOperator(task_id="start", dag=dag)
end = DummyOperator(task_id="end", dag=dag)

start >> end
    